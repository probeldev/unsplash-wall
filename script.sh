set -x

timestamp=$(date +%s)
echo $timestamp

mkdir /tmp/wall || echo 2

curl https://unsplash.com/ | xq | grep "\"download\":" | sed 's#<script id="dehydrated-data" type="application/json">##g' | sed 's#</script>##g' | jq | grep "\"download\":" | sed 's# ##g' | sed "s#\"download#wget -O /tmp/wall/wall-$timestamp.jpg #g" | sed 's#"##g' | sed 's#:##g' | sed 's#https#https:#g' | sort -R | head -n 1 >/tmp/wall/download.sh &&
	rm -f /tmp/wall/*.jpg

# curl https://unsplash.com/ | xq | grep "\"download\":" | sed 's#<script id="dehydrated-data" type="application/json">##g' | sed 's#</script>##g' | jq | grep "\"download\":" | sed 's# ##g' | sed "s#\"download#wget -O wall.jpg #g" | sed 's#"##g' | sed 's#:##g' | sed 's#https#https:#g' | sort -R | head -n 1 >download.sh &&
sh /tmp/wall/download.sh &&
	# killall hyprpaper &&
	# nohup hyprpaper &
	rm -f /tmp/wall/setbg.sh &&
	echo "hyprctl hyprpaper preload \"/tmp/wall/wall-$timestamp.jpg\"" >>/tmp/wall/setbg.sh &&
	echo "hyprctl hyprpaper wallpaper \",/tmp/wall/wall-$timestamp.jpg\"" >>/tmp/wall/setbg.sh &&
	sh /tmp/wall/setbg.sh
