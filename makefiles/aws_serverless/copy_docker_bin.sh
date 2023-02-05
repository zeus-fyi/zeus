#!/bin/sh

#docker run -d zeusfyi/ethereumbls:latest /usr/bin/ethereumsignbls

#zip main.zip main
#
#        docker run -d ${{ env.SERVERLESS_REPO }}/ethereumbls:latest
#
#        docker run -d zeusfyi/ethereumbls:latest
#        docker cp $(docker container ls | awk 'NR==2 {print $1}'):/usr/bin/ethereumsignbls main
#docker run -it zeusfyi/ethereumbls:latest /usr/bin/ethereumsignbls bash -c "echo '/usr/bin/ethereumsignbls > ."
#
#	docker run -it --entrypoint /bin/bash ${SERVERLESS_REPO}/ethereumbls:latest
#
#docker run --entrypoint "/bin/bash -c 'sleep 1; /bin/bash'" zeusfyi/ethereumbls:latest && docker cp $(docker container ls | awk 'NR==2 {print $1}'):/usr/bin/ethereumsignbls main
#
#docker run --entrypoint "/bin/bash -c 'sleep 1; /bin/bash'" zeusfyi/ethereumbls:latest
#
#docker run --entrypoint "-c 'sleep 1" zeusfyi/ethereumbls:latest

docker run -d zeusfyi/ethereumbls:latest sleep 5
docker cp $(docker container ls | awk 'NR==2 {print $1}'):/usr/bin/ethereumsignbls main
docker container stop $(docker container ls | awk 'NR==2 {print $1}')