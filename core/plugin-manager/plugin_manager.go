package pluginmanager

// PluginManager defines methods for managing hooks and middleware injection.
type PluginManager interface {
	RegisterPreHook(stepID string, hookFunc interface{}) error
	RegisterPostHook(stepID string, hookFunc interface{}) error
	InjectMiddleware(middlewareFunc interface{}) error
}
