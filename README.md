# gosshtool
similar to salt-ssh tools,  类似于salt-ssh的一个GO实现远程服务器管理的工具, 具有web页面操作,在多远程主机操作时，使用goroutines多协程后，ssh响应速度比ansible,salt-ssh快几倍</br>
gosshtool:</br>
![image](https://github.com/kggg/gosshtool/blob/master/static/img/gosshtool.png)</br>
ansible:</br>
![image](https://github.com/kggg/gosshtool/blob/master/static/img/ansible.png)
</br>
## 数据存储方式
   mysql

## 命令行模式
Usage: gossh host [host|group] options [cmd|copyfile]</br>
	  -h : specified a remote host, use , split one or more  host </br>
	  -g : specified a remote hostgroup </br>
	  -E : Regrex match a remote host name default </br>
	  -m : select a module, -m [cmd|copy]> </br>
		   copy : [src, dest,mode,force,backup,user,owner]</br></br>

e.g.:   gossh -h steven -m cmd 'uptime'</br>
        gossh -h dbserver -m copy "src=/etc/nginx/nginx.conf dest=/etc/nginx/nginx.conf mode=0644"</br>

## 主机列表
![image](https://github.com/kggg/gosshtool/blob/master/static/img/Screenshot-host.png)

## ssh远程操作
![image](https://github.com/kggg/gosshtool/blob/master/static/img/Screenshot-command.png)

## 群组ssh远程操作

## 配置文件管理


## 用户管理


## 用户权限 

