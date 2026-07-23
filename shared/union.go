// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsPermissionPatternUnion() {}

type UnionBool bool

func (UnionBool) ImplementsConfigProviderOptionsTimeoutUnion()  {}
func (UnionBool) ImplementsMcpAddParamsConfigRemoteOAuthUnion() {}

type UnionInt int64

func (UnionInt) ImplementsConfigProviderOptionsTimeoutUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsConfigProviderOptionsTimeoutUnion() {}
