# ./wrk -D exp -t 4 -c 8 -d 30s -L -s ./scripts/social-network/compose-post.lua http://10.0.0.14:31000/wrk2-api/post/compose -R 200
./wrk -D exp -t 4 -c 8 -d 30s -L -s ./scripts/social-network/compose-post.lua http://sn-nginx.test.com/wrk2-api/post/compose -R 200
# ./wrk -D exp -t 16 -c 32 -d 30s -L -s ./scripts/social-network/compose-post.lua http://sn-nginx.test.com/wrk2-api/post/compose -R 2000