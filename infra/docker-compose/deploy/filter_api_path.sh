# /bin/bash
cd `dirname $0`
for target_path in `ls ./../../../app/gin/ | grep -v "cmd" | grep -xv "main"`
do
    echo "- ../../../app/gin/$target_path:/usr/src/app/$target_path"
done
