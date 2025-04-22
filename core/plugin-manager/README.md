Responsibilities:
*	Handle pre/post hooks
*	Enable injecting middleware / custom behavior

```golang
type PluginManager struct {
    PreHooks  map[string]func()
    PostHooks map[string]func()
}

func (pm *PluginManager) RegisterPreHook(name string, hook func())
func (pm *PluginManager) RegisterPostHook(name string, hook func())
```