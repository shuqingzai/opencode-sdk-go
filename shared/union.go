// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

type UnionString string

func (UnionString) ImplementsPermissionPatternUnion()      {}
func (UnionString) ImplementsConfigV2ReferenceUnionParam() {}

type UnionBool bool

func (UnionBool) ImplementsConfigProviderOptionsTimeoutUnion()         {}
func (UnionBool) ImplementsMcpAddParamsConfigRemoteOAuthUnion()        {}
func (UnionBool) ImplementsProviderModelCapabilitiesInterleavedUnion() {}

type UnionInt int64

func (UnionInt) ImplementsConfigProviderOptionsTimeoutUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsConfigProviderOptionsTimeoutUnion() {}
