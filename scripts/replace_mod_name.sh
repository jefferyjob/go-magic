#!/bin/bash

# 指定要替换的包地址
old_package="github.com/jefferyjob/go-magic"

# 提示用户输入新的包地址
echo "请输入新的包地址："
read -r new_package

# 显示用户输入的新包地址并要求确认
echo "新的包地址为：$new_package"
echo "确认要将旧的包地址 $old_package 替换为新的包地址 $new_package 吗？(y/n)"
read -r confirm

# 判断用户的确认选项
if [ "$confirm" != "y" ]; then
    echo "取消替换操作。"
    exit 0
fi

# 指定要遍历的目录
directory="../"

# 递归遍历目录下的所有Go文件
find "$directory" -type f -name "*.go" | while read -r file; do
    # 替换文件中的包地址
    sed -i "s|$old_package|$new_package|g" "$file"
done

# 递归遍历目录下的所有Go文件和.mod文件
#find "$directory" \( -name "*.go" -o -name "*.mod" \) | while read -r file; do
#    # 替换文件中的包地址
#    sed -i "s|$old_package|$new_package|g" "$file"
#done

echo "替换完成！"
