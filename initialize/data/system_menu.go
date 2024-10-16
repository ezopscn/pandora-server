package data

import (
	"errors"
	"gorm.io/gorm"
	"pandora-server/global"
	"pandora-server/model"
)

// 菜单数据
var systemMenus = []model.SystemMenu{
	{
		Id:          1100,
		Label:       "工作空间",
		Icon:        "DesktopOutlined",
		Key:         "/dashboard",
		Sort:        1100,
		ParentId:    0,
		SystemRoles: systemRoles,
	},
	{
		Id:       1200,
		Label:    "集群列表",
		Icon:     "PartitionOutlined",
		Key:      "/clusters",
		Sort:     1200,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
	},
	{
		Id:       1300,
		Label:    "集群管理",
		Icon:     "KubernetesOutlined",
		Key:      "/cluster",
		Sort:     1300,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
		Children: []model.SystemMenu{
			{
				Id:       1310,
				Label:    "集群概览",
				Key:      "/cluster/overview",
				Sort:     1310,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1320,
				Label:    "节点管理（Node）",
				Key:      "/cluster/node",
				Sort:     1320,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1330,
				Label:    "名称空间（Namespace）",
				Key:      "/cluster/namespace",
				Sort:     1330,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1340,
				Label:    "角色（Role）",
				Key:      "/cluster/namespace-role",
				Sort:     1340,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1350,
				Label:    "绑定（Role Binding）",
				Key:      "/cluster/namespace-role-binding",
				Sort:     1350,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1360,
				Label:    "集群角色（Cluster Role）",
				Key:      "/cluster/cluster-role",
				Sort:     1360,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1370,
				Label:    "集群绑定（Cluster Role Binding）",
				Key:      "/cluster/cluster-role-binding",
				Sort:     1370,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1380,
				Label:    "服务账号（Service Account）",
				Key:      "/cluster/service-account",
				Sort:     1380,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1390,
				Label:    "集群证书（Certificate）",
				Key:      "/cluster/certificate",
				Sort:     1390,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
		},
	},
	{
		Id:       1400,
		Label:    "工作负载",
		Icon:     "DeploymentUnitOutlined",
		Key:      "/workload",
		Sort:     1400,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
		},
		Children: []model.SystemMenu{
			{
				Id:       1410,
				Label:    "Pod",
				Key:      "/workload/pod",
				Sort:     1410,
				ParentId: 1400,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1420,
				Label:    "副本集（Replica Set）",
				Key:      "/workload/replica-set",
				Sort:     1420,
				ParentId: 1400,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1430,
				Label:    "部署（Deployment）",
				Key:      "/workload/deployment",
				Sort:     1430,
				ParentId: 1400,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1440,
				Label:    "守护进程集（Daemon Set）",
				Key:      "/workload/daemon-set",
				Sort:     1440,
				ParentId: 1400,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1450,
				Label:    "有状态集（Stateful Set）",
				Key:      "/workload/stateful-set",
				Sort:     1450,
				ParentId: 1400,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1460,
				Label:    "普通任务（Job）",
				Key:      "/workload/job",
				Sort:     1460,
				ParentId: 1400,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1470,
				Label:    "定时任务（Cron Job）",
				Key:      "/workload/cron-job",
				Sort:     1470,
				ParentId: 1400,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
		},
	},
	{
		Id:       1500,
		Label:    "服务发现",
		Icon:     "ApiOutlined",
		Key:      "/service",
		Sort:     1500,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
		},
		Children: []model.SystemMenu{
			{
				Id:       1510,
				Label:    "服务（Service）",
				Key:      "/service/svc",
				Sort:     1510,
				ParentId: 1500,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1520,
				Label:    "流量入口类（Ingress Class）",
				Key:      "/service/ingress-class",
				Sort:     1520,
				ParentId: 1500,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1530,
				Label:    "流量入口（Ingress）",
				Key:      "/service/ingress",
				Sort:     1530,
				ParentId: 1500,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
		},
	},
	{
		Id:       1600,
		Label:    "存储管理",
		Icon:     "HddOutlined",
		Key:      "/storage",
		Sort:     1600,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
		},
		Children: []model.SystemMenu{
			{
				Id:       1610,
				Label:    "存储类（Storage Class）",
				Key:      "/storage/class",
				Sort:     1610,
				ParentId: 1600,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1620,
				Label:    "持久卷（Persistent Volume）",
				Key:      "/storage/persistent-volume",
				Sort:     1620,
				ParentId: 1600,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1630,
				Label:    "持久卷申领（Persistent Volume Claim）",
				Key:      "/storage/persistent-volume-claim",
				Sort:     1630,
				ParentId: 1600,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
		},
	},
	{
		Id:       1700,
		Label:    "配置管理",
		Icon:     "FileProtectOutlined",
		Key:      "/config",
		Sort:     1700,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
		},
		Children: []model.SystemMenu{
			{
				Id:       1710,
				Label:    "普通配置（Config Map）",
				Key:      "/config/config-map",
				Sort:     1710,
				ParentId: 1700,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1720,
				Label:    "敏感配置（Secret）",
				Key:      "/config/secret",
				Sort:     1720,
				ParentId: 1700,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
		},
	},
	{
		Id:       1800,
		Label:    "资源设置",
		Icon:     "SolutionOutlined",
		Key:      "/resource",
		Sort:     1800,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
		Children: []model.SystemMenu{
			{
				Id:       1810,
				Label:    "资源组别",
				Key:      "/resource/group",
				Sort:     1810,
				ParentId: 1800,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1820,
				Label:    "资源授权",
				Key:      "/resource/permission",
				Sort:     1820,
				ParentId: 1800,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
		},
	},
	{
		Id:       1900,
		Label:    "系统设置",
		Icon:     "SettingOutlined",
		Key:      "/system",
		Sort:     1900,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
		},
		Children: []model.SystemMenu{
			{
				Id:       1910,
				Label:    "用户中心",
				Key:      "/system/user",
				Sort:     1910,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
				},
			},
			{
				Id:       1920,
				Label:    "用户角色",
				Key:      "/system/role",
				Sort:     1920,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
				},
			},
			{
				Id:       1930,
				Label:    "系统菜单",
				Key:      "/system/menu",
				Sort:     1930,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
				},
			},
			{
				Id:       1940,
				Label:    "系统接口",
				Key:      "/system/api",
				Sort:     1940,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
				},
			},
			{
				Id:       1950,
				Label:    "系统授权",
				Key:      "/system/permission",
				Sort:     1950,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
				},
			},
			{
				Id:       1960,
				Label:    "系统设置",
				Key:      "/system/setting",
				Sort:     1960,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
				},
			},
		},
	},
	{
		Id:          2000,
		Label:       "个人中心",
		Icon:        "UsergroupAddOutlined",
		Key:         "/user-center",
		Sort:        2000,
		ParentId:    0,
		SystemRoles: systemRoles,
	},
	{
		Id:          2100,
		Label:       "消息通知",
		Icon:        "BellOutlined",
		Key:         "/message",
		Sort:        2100,
		ParentId:    0,
		SystemRoles: systemRoles,
	},
	{
		Id:          9999,
		Label:       "获取帮助",
		Icon:        "QuestionCircleOutlined",
		Key:         "/help",
		Sort:        9999,
		ParentId:    0,
		SystemRoles: systemRoles,
	},
}

// 递归插入数据方法
func insertMenusData(menus []model.SystemMenu) {
	var m model.SystemMenu
	for _, item := range menus {
		// 查看数据是否存在，如果不存在才执行创建，注意 Key 是关键字，需要特别处理，否则会报错
		err := global.MySQLDB.Where("id = ? OR label = ? OR `key` = ?", item.Id, item.Label, item.Key).First(&m).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.MySQLDB.Create(&item)
		}
		// 递归插入子菜单
		if len(item.Children) > 0 {
			insertMenusData(item.Children)
		}
	}
}

// 菜单数据初始化
func InitializeSystemMenu() {
	insertMenusData(systemMenus)
}
