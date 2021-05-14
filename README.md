# Step 1: Installing the Software
Download the test movie file  "https://download.blender.org/peach/bigbuckbunny_movies/"
#
Download ffmpeg from "https://www.videohelp.com/software?d=ffmpeg-4.4-full_build.7z"
#
Download the latest installer of GPAC from "https://gpac.wp.imt.fr/downloads/gpac-nightly-builds/"
# 
Extract the downloaded ffmpeg zip file to "c:\ffmpeg"
#
Navigate to the "bin" folder under c:\ffmpeg and copy the address using Ctrl+C
# Step 2: Testing the Programs
1. ffmpeg
2. mp4box
# Step 3: Creating Your Workspace
# Step 4: Opening the Command Line
# Step 5: Encoding the Video
1. ffmpeg -i input.avi -s 160x90 -c:v libx264 -b:v 250k -g 90 -an input_video_160x90_250k.mp4
2. ffmpeg -i input.avi -s 320x180 -c:v libx264 -b:v 500k -g 90 -an input_video_320x180_500k.mp4
3. ffmpeg -i input.avi -s 640x360 -c:v libx264 -b:v 750k -g 90 -an input_video_640x360_750k.mp4
4. ffmpeg -i input.avi -s 640x360 -c:v libx264 -b:v 1000k -g 90 -an input_video_640x360_1000k.mp4
5. ffmpeg -i input.avi -s 1280x720 -c:v libx264 -b:v 1500k -g 90 -an input_video_1280x720_1500k.mp4


* ffmpeg -i input.avi -vn -acodec libvorbis -ab 128k -dash 1 my_audio.webm
* ffmpeg -i in.video -c:v libvpx-vp9 -keyint_min 150 -g 150 -tile-columns 4 -frame-parallel 1  -f webm -dash 1 \
-an -vf scale=160:90 -b:v 250k -dash 1 video_160x90_250k.webm

* ffmpeg -i input.avi -c:v libvpx-vp9 -keyint_min 150 -g 150 -tile-columns 4 -frame-parallel 1  -f webm -dash 1 \
-an -vf scale=320:180 -b:v 500k -dash 1 video_320x180_500k.webm

* ffmpeg -i input.avi-c:v libvpx-vp9 -keyint_min 150 -g 150 -tile-columns 4 -frame-parallel 1  -f webm -dash 1 \
-an -vf scale=640:360 -b:v 750k -dash 1 video_640x360_750k.webm

* ffmpeg -i input.avi -c:v libvpx-vp9 -keyint_min 150 -g 150 -tile-columns 4 -frame-parallel 1  -f webm -dash 1 \
-an -vf scale=640:360 -b:v 1000k -dash 1 video_640x360_1000k.webm

* ffmpeg -i input.avi -c:v libvpx-vp9 -keyint_min 150 -g 150 -tile-columns 4 -frame-parallel 1  -f webm -dash 1 \
-an -vf scale=1280:720 -b:v 1500k -dash 1 video_1280x720_1500k.webm

All file
* ffmpeg -i input.avi -c:v libvpx-vp9 -keyint_min 150 \
-g 150 -tile-columns 4 -frame-parallel 1  -f webm -dash 1 \
-an -vf scale=160:90 -b:v 250k -dash 1 video_160x90_250k.webm \
-an -vf scale=320:180 -b:v 500k -dash 1 video_320x180_500k.webm \
-an -vf scale=640:360 -b:v 750k -dash 1 video_640x360_750k.webm \
-an -vf scale=640:360 -b:v 1000k -dash 1 video_640x360_1000k.webm \
-an -vf scale=1280:720 -b:v 1500k -dash 1 video_1280x720_1500k.webm

# Create the manifest file:
ffmpeg \
  -f webm_dash_manifest -i video_160x90_250k.webm \
  -f webm_dash_manifest -i video_320x180_500k.webm \
  -f webm_dash_manifest -i video_640x360_750k.webm \
  -f webm_dash_manifest -i video_1280x720_1500k.webm \
  -f webm_dash_manifest -i my_audio.webm \
  -c copy \
  -map 0 -map 1 -map 2 -map 3 -map 4 \
  -f webm_dash_manifest \
  -adaptation_sets "id=0,streams=0,1,2,3 id=1,streams=4" \
  manifest.mpd


# Step 6: Encode the Audio
ffmpeg -i input.avi -c:a aac -b:a 128k -vn input_audio_128k.mp4
# Step 7: Dashifying the Encoded Files
mp4box -dash 5000 -rap -profile dashavc264:onDemand -mpd-title BBB -out manifest.mpd -frag 2000 input_audio_128k.mp4 input_video_160x90_250k.mp4 input_video_320x180_500k.mp4 input_video_640x360_750k.mp4 input_video_640x360_1000k.mp4 input_video_1280x720_1500k.mp4
# Step 8: Setting Up the Web Server using golang
# Step 9: Enabling CORS using golang
# Step 10: Adding the DASH MIME Type using golang
# Step 11: Setting Up Your Network using golang
# Step 12: Streaming the Video
VLC media Player
Go to Media-> Open Location from clipboard-> Then enter url-> play
Successfully play your video
