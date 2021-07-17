<h1 align="center">μRTMP</h1>

<p align="center">Simple low-overhead minimal RTMP server with embedded UI to watch your live streams online.</p>

## Usage

Simplest way to run `μRTMP` server is using `Docker`:

```sh
docker run -it --rm --name urtmp -p 1935:1935 -p 8080:8080 jkuri/urtmp
```

Then open your browser at `http://localhost:8080` where you can watch your published streams.

Example of publishing the stream using `ffmpeg` from MacOS:

```sh
ffmpeg -f avfoundation -video_size 800x600 -framerate 24 -i 0 -vcodec libx264 -preset ultrafast -tune zerolatency -f flv "rtmp://localhost/live/stream"
```

If stream is working it should be available online at `http://localhost:8080`:

<p align="center">
  <img src="https://user-images.githubusercontent.com/1796022/126041179-b1cd220d-b4ed-4b6d-80ee-214067bf6ae0.png" alt="uRTMP live stream">
</p>


## Compile from source

This step assumes that you already have `Node.JS` and `Golang` installed on your system.

```sh
git clone https://github.com/jkuri/urtmp
cd urtmp
make install
make
```

If everything went well, you should have build artifacts in `build/` folder.

## License

MIT
