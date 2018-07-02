import os
import os.path
path =os.getcwd()
listdir =os.listdir(path)
html =[]
for dirs in listdir:
    if (dirs[-1]=="l") and (dirs[0].isnumeric()):
        html.append(dirs)

html.sort()
print(html)
k=0
for item in html:
    if 0<k<len(html)-1:
        next_path = html[k+1]
        l_path =html[k-1]
    elif k==0:
        l_path =html[-1]
        next_path=html[1]
    else:
        l_path =html[-2]
        next_path=html[0]
        k=0
    with open(item,"a") as f:
       f.write('\n<a href="./'+next_path+'">下一页</a>\n')
       f.write('\n<a href="./'+l_path+'">上一页</a>\n')

    k = k+1
print("end!")
