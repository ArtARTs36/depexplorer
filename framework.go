package depexplorer

type FrameworkName string

const (
	FrameworkNameGin   FrameworkName = "Gin Web Framework"
	FrameworkNameFiber               = "Fiber"

	FrameworkNameSymfony FrameworkName = "Symfony"
	FrameworkNameLaravel FrameworkName = "Laravel"

	FrameworkNameVueJS FrameworkName = "Vue.js"
	FrameworkNameReact FrameworkName = "React"
)

type Framework struct {
	Name    FrameworkName
	Version Version
}

var frameworksDepMap = map[DependencyManager]map[string]FrameworkName{
	DependencyManagerNone: {},
	DependencyManagerGoMod: {
		"github.com/gin-gonic/gin":    FrameworkNameGin,
		"github.com/gofiber/fiber/v2": FrameworkNameFiber,
	},
	DependencyManagerComposer: {
		"symfony/framework-bundle": FrameworkNameSymfony,
		"laravel/framework":        FrameworkNameLaravel,
	},
	DependencyManagerNPM: {
		"vue":   FrameworkNameVueJS,
		"react": FrameworkNameReact,
	},
}

func dependencyToFramework(depManager DependencyManager, dep *Dependency) (*Framework, bool) {
	deps, ok := frameworksDepMap[depManager]
	if !ok {
		return nil, false
	}

	frName, ok := deps[dep.Name]
	if !ok {
		return nil, false
	}

	return &Framework{
		Name:    frName,
		Version: dep.Version,
	}, true
}
