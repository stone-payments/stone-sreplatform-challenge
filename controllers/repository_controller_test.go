package controllers

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	repositoriesv1alpha1 "github.com/stone-payments/stone-sreplatform-challenge/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Repository Controller", func() {
	BeforeEach(func() {})

	AfterEach(func() {})

	Context("The Repository controller", func() {
		It("Should ...", func() {
			secretRef := repositoriesv1alpha1.SecretKeyReference{
				Name:      "mock",
				Namespace: "default",
				Key:       "token",
			}
			err := createReferencedSecret(secretRef.Name, secretRef.Namespace, secretRef.Key, "MYF4K3P4T")
			Expect(err).ToNot(HaveOccurred())

			nn := types.NamespacedName{
				Name:      "test-repository",
				Namespace: "default",
			}

			_ = &repositoriesv1alpha1.Repository{
				ObjectMeta: metav1.ObjectMeta{
					Name:      nn.Name,
					Namespace: nn.Namespace,
				},
				Spec: repositoriesv1alpha1.RepositorySpec{
					Name:           "sample",
					Owner:          "sample",
					Type:           "OpenSource",
					CredentialsRef: secretRef,
				},
			}

			By("...")
			// TODO
		})
	})
})

func createReferencedSecret(name, namespace, key, token string) error {
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: map[string][]byte{
			key: []byte(token),
		},
	}
	return k8sClient.Create(context.Background(), secret)
}
