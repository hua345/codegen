# base image
FROM centos

# MAINTAINER
MAINTAINER chenjianhua 2290910211@qq.com
# update yum
RUN yum -y install wget
RUN wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
RUN yum clean all && yum makecache && yum -y update

# update epel
RUN wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo
RUN yum repolist
RUN yum install -y gcc gcc-c++ glibc make autoconf git
# Commands when creating a new container
CMD ["cat","/etc/redhat-release"]