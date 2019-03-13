#安装gin
go get -u github.com/gin-gonic/gin
#安装mysql驱动
go get -u github.com/go-sql-driver/mysql
#安装gorm
go get -u github.com/jinzhu/gorm
#安装yaml


#docker 安装mysql
docker pull mysql:5.6
docker run --name mysql -v /root/mysql:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.6