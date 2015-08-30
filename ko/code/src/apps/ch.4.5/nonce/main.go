// A nonce is a number or string used only once.
// This is useful for generating a unique token for login pages to prevent duplicate submissions.
package nonce

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

// Contains a unique token
type Nonce struct {
	Token string
}

// Keeps track of marked/used tokens
type Nonces struct {
	hashs map[string]bool
}

func New() Nonces {
	return Nonces{make(map[string]bool)}
}
func (n *Nonces) NewNonce() Nonce {
	return Nonce{n.NewToken()}
}

// Returns a new unique token
func (n *Nonces) NewToken() string {
	t := createToken()
	for n.HasToken(t) {
		t = createToken()
	}
	return t
}

// Checks if token has been marked.
func (n *Nonces) HasToken(token string) bool {
	return n.hashs[token] == true
}
func (n *Nonces) MarkToken(token string) {
	n.hashs[token] = true
}
func (n *Nonces) CheckToken(token string) error {
	if token == "" {
		return errors.New("No token supplied")
	}
	if n.HasToken(token) {
		return errors.New("Duplicate submission.")
	}
	return nil
}
func (n *Nonces) CheckThenMarkToken(token string) error {
	defer n.MarkToken(token)
	if err := n.CheckToken(token); err != nil {
		return err
	}
	return nil
}
func createToken() string {
	h := md5.New()
	now := time.Now().Unix()
	io.WriteString(h, strconv.FormatInt(now, 10))
	io.WriteString(h, strconv.FormatInt(rand.Int63(), 10))
	return fmt.Sprintf("%x", h.Sum(nil))
}
