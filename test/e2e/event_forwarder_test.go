package e2e_test

import (
	"net/http"
	"os"
	"syscall"
	"time"

	api "github.com/appscode/kubed/apis/kubed/v1alpha1"
	"github.com/appscode/kubed/test/e2e/framework"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
)

var _ = Describe("Event-forwarder", func() {
	var (
		f              *framework.Invocation
		pvc            *core.PersistentVolumeClaim
		pod            *core.Pod
		clusterConfig  api.ClusterConfig
		stopCh         chan struct{}
		stopServer     chan os.Signal
		requests       []*http.Request
		notifierSecret *core.Secret
	)

	BeforeEach(func() {
		f = root.Invoke()
	})

	JustBeforeEach(func() {
		if missing, _ := BeZero().Match(notifierSecret); missing {
			Skip("Missing notifier-secret")
		}

		if f.SelfHostedOperator {
			pod, err := f.OperatorPod()
			Expect(err).NotTo(HaveOccurred())

			if pod.Spec.NodeName != "minikube" {
				Skip("Skipping test. Reason: This test is designed only for minikube.")
			}

			By("Restarting kubed operator")
			err = f.RestartKubedOperator(&clusterConfig)
			Expect(err).NotTo(HaveOccurred())
		} else {
			By("Starting Kubed")
			stopCh = make(chan struct{})
			err := f.RunKubed(stopCh, clusterConfig)
			Expect(err).NotTo(HaveOccurred())

			By("Waiting for API server to be ready")
			root.EventuallyAPIServerReady().Should(Succeed())
			time.Sleep(time.Second * 5)
		}
	})

	AfterEach(func() {
		if !f.SelfHostedOperator {
			close(stopCh)
		}

		err := f.DeleteSecret(notifierSecret.ObjectMeta)
		Expect(err).NotTo(HaveOccurred())

		err = f.WaitUntilSecretDeleted(notifierSecret.ObjectMeta)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Event-forwarder for Webhook Server", func() {
		svc := &core.Service{}
		ep := &core.Endpoints{}

		BeforeEach(func() {
			requests = make([]*http.Request, 0)
			stopServer = make(chan os.Signal, 1)

			By("Starting Webhook Server")
			f.RunWebhookServer(stopServer, &requests)

			notifierSecret = f.SecretForWebhookNotifier()

			if f.SelfHostedOperator {
				svc = f.ServiceForWebhook()
				ep = f.LocalEndPointForWebhook(svc)

				_, err := f.KubeClient.CoreV1().Services(svc.Namespace).Create(svc)
				Expect(err).NotTo(HaveOccurred())

				_, err = f.KubeClient.CoreV1().Endpoints(ep.Namespace).Create(ep)
				Expect(err).NotTo(HaveOccurred())

				notifierSecret.Data["WEBHOOK_URL"] = []byte("http://" + svc.Name + "." + svc.Namespace + ".svc")
			}

			By("Creating notifier secret: " + notifierSecret.Name)
			_, err := f.CreateSecret(notifierSecret)
			Expect(err).NotTo(HaveOccurred())

			By("Creating clusterConfiguration")
			clusterConfig = f.EventForwarderClusterConfig()
			clusterConfig.NotifierSecretName = notifierSecret.Name
			clusterConfig.EventForwarder.Receivers = framework.WebhookReceiver()
		})

		AfterEach(func() {
			By("Closing Webhook Server")
			stopServer <- os.Signal(syscall.SIGINT)
			defer close(stopServer)

			if f.SelfHostedOperator {
				err := f.DeleteService(svc.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())
			}
		})

		Context("PVC add eventer", func() {

			BeforeEach(func() {
				pvc = f.NewPersistentVolumeClaim()
			})

			AfterEach(func() {
				err := f.DeletePersistentVolumeClaim(pvc.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should send notification to Webhook Server", func() {

				By("Creating PVC: " + pvc.Name)
				err := f.CreatePersistentVolumeClaim(*pvc)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting for notification in Webhook Server")
				f.EventuallyNotifiedToWebhookServer(&requests, "PersistentVolumeClaim").Should(BeTrue())
			})
		})

		Context("Pod Warning Event", func() {

			BeforeEach(func() {
				pod = f.NewPod()
			})

			AfterEach(func() {
				err := f.DeletePod(pod.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				err = f.WaitUntilPodTerminated(pod.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())
			})

			It("Check warning event", func() {

				By("Creating pod: " + pod.Name)
				_, err := f.CreatePod(pod)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting for notification in Webhook Server")
				f.EventuallyNotifiedToWebhookServer(&requests, pod.Name).Should(BeTrue())
			})
		})
	})
})
