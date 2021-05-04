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
2. ffmpeg -i input.avi -s 160x90 -c:v libx264 -b:v 250k -g 90 -an input_video_160x90_250k.mp4
3. ffmpeg -i input.avi -s 320x180 -c:v libx264 -b:v 500k -g 90 -an input_video_320x180_500k.mp4
4. ffmpeg -i input.avi -s 640x360 -c:v libx264 -b:v 750k -g 90 -an input_video_640x360_750k.mp4
5. ffmpeg -i input.avi -s 640x360 -c:v libx264 -b:v 1000k -g 90 -an input_video_640x360_1000k.mp4
6. ffmpeg -i input.avi -s 1280x720 -c:v libx264 -b:v 1500k -g 90 -an input_video_1280x720_1500k.mp4

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
#
Go to Media-> Open Location from clipboard-> Then enter url-> play
#
Successfully play your video
