###数据库设计

1.检测表(checker_server)
* id int(10)
* name varchar(255) not null
* check_image varchar(255) not null
* step int(2) default 1
* script varchar(255) not null
* ip_lists varchar(255) not null

---

	字段说明:
		name: 设置检测名称
		check_image: 选择执行检测的镜像
		step: 设置脚本运行步长(暂定以分钟解析)
		script: 添加所执行的cheack脚本文件
		ip_lists :设置需要检测的ip列表(,分隔)
		
		
	运行说明:
		所有靶机容器启动后,设置检测模板各参数,成功后保存到数据库.
		相应的可以获取已执行的检测.	

		
	


	