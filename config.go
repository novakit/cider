package cider

type Config struct {
	// Service 被托管应用提供服务的配置
	Service struct {
		// Command 被托管应用的启动命令
		Command []string
		// Port 被托管应用的 HTTP 服务监听地址
		// 默认使用随机端口号
		// 会以 'CIDER_SERVICE_PORT' 环境变量提供给被托管应用
		Port int
	}
	// Inlet Cider 对外提供服务的配置
	Inlet struct {
		// Port 对外提供 HTTP 服务监听的地址
		// 默认使用 3000 端口
		// 会使用 'CIDER_INLET_PORT' 环境变量提供给被托管应用
		// 一般情况下，被托管应用不需要关心该端口号
		Port int
	}
	// Outlet 依赖服务的配置
	Outlet struct {
		// Port 依赖服务聚合端口
		// 所有的依赖服务会聚合在这个端口
		// 默认使用随机端口号
		// 会以 'CIDER_OUTLET_PORT' 环境变量提供给被托管应用
		Port int
		// Sources 依赖服务
		Sources []struct {
			// Addr 依赖服务地址
			// 支持 tcp://IP:PORT
			// 支持 tcp+k8s://service.namespace:PORT
			Addr string
			// Path 依赖服务 HTTP 路径前缀
			// 用以将不同的服务聚合在同一个端口
			Path string
		}
	}
}
