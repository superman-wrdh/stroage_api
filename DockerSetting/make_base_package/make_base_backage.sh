#!/bin/sh
mkdir temp
echo "开始制作镜像..."
image_tag=`date +%Y%m%d` #_%H%M
echo "当前时间：$image_tag"
docker build -t ccr.ccs.tencentyun.com/66super/go_base:1.12.7-v${image_tag} .
echo "制作镜像成功!"