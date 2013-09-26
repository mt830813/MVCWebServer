package Common

type decorateRegistcontext struct {
	currentContext *registContext
	nextContext    *decorateRegistcontext
}
