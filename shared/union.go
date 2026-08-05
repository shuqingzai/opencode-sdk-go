// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionBool bool

func (UnionBool) ImplementsConfigProviderOptionsTimeoutUnion() {}
func (UnionBool) ImplementsMcpOAuthConfigUnion()               {}

type UnionInt int64

func (UnionInt) ImplementsConfigProviderOptionsTimeoutUnion() {}
