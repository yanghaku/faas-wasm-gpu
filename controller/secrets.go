package controller

import "github.com/openfaas/faas-provider/types"

// SecretsClient exposes the standardized CRUD behaviors for secrets.
type SecretsClient interface {
	// List returns a list of available function secrets.
	List() (names []*types.Secret, err error)
	// Create adds a new secret
	Create(secret types.Secret) error
	// Replace updates the value of a function secret
	Replace(secret types.Secret) error
	// Delete removes a function secret
	Delete(name string) error
	// GetSecrets return a secret map for the specific name list
	GetSecrets(secretsNames []string) (map[string]*types.Secret, error)
}

// DefaultMemorySecretsClient the default SecretsClient implement, just use a map in memory
type DefaultMemorySecretsClient struct {
	// secrets maintain a map name -> secret in memory
	secrets map[string]*types.Secret
}

// List implements the list function
func (c *DefaultMemorySecretsClient) List() (names []*types.Secret, err error) {
	res := make([]*types.Secret, len(c.secrets))
	idx := 0
	for _, s := range c.secrets {
		res[idx] = s
		idx += 1
	}
	return res, nil
}

// Create add a new secret
func (c *DefaultMemorySecretsClient) Create(secret types.Secret) error {
	c.secrets[secret.Name] = &secret
	return nil
}

// Replace updates the value of a function secret
func (c *DefaultMemorySecretsClient) Replace(secret types.Secret) error {
	c.secrets[secret.Name] = &secret
	return nil
}

// Delete removes a function secret
func (c *DefaultMemorySecretsClient) Delete(name string) error {
	delete(c.secrets, name)
	return nil
}

// GetSecrets return a secret map for the specific name list
func (c DefaultMemorySecretsClient) GetSecrets(secretsNames []string) (map[string]*types.Secret, error) {
	res := map[string]*types.Secret{}
	for _, v := range secretsNames {
		s, ok := c.secrets[v]
		if ok {
			res[v] = s
		}
	}
	return res, nil
}

// NewDefaultMemorySecretsClient return a new handler
func NewDefaultMemorySecretsClient() *DefaultMemorySecretsClient {
	return &DefaultMemorySecretsClient{
		secrets: map[string]*types.Secret{},
	}
}
