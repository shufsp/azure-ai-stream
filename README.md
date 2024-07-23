### superbarosa's AI superplex
This is a monolothic repository for superbarosa's AI streams! Any AI stream I do has something from this repo being used.
This repository is meant to be ran on linux, so you may have a lot of trouble if you're a windows or mac user.


## Modules
``barosa-azure`` - this is the folder for Azure computer vision. It's vaguely named at the moment..

``barosa-integration-testing`` - folder reserved for integration tests. you can freely ignore htis

``barosa-microphone`` - folder for Azure audio AI processing, such as transcript and AI voice

``barosa-screen-capture`` - folder for window screenshotting and image processing that is intended to be send to Azure computer vision 

## Prerequisites
- You need api keys from Azure (you have to provide payment info, so it may cost)
- You need ``maim`` and ``xdotool`` installed and in ``$PATH``
- You need GO installed, since the entire backend is made in GO
- You need `node` installed with `npm`, since the "web server" hosting the front-end is a barebones `express` server

## Setup 
This is the generic setup for using the barosa computer vision modules. Parts of this setup you can skip will be labeled ``[OPTIONAL]``

### Azure Environment Variables
You will need to set up the environment variables for _your_ Azure computer vision resource. *You can't use mine, I'm poor*. There's a free usage threshold, but you still have to provide payment info when you sign up. 
Basically, just be prepared to be charged for using your Azure keys in this project.

Open the `barosa-azure` folder

Then, create a `.env` file and fill the environment values below in it:
```env
VISION_KEY=replace_with_your_azure_computer_vision_api_key
VISION_ENDPOINT=replace_with_your_azure_computer_vision_endpoint
VISION_REGION=replace_with_your_azure_resource_region
```
<br />

Check that you've set up everything correctly by running the `main.go` file inside of `barosa-azure`. You can set the image filename as the example Krita screen capture included in this repository
```bash
go run main.go "../barosa-screen-capture/Krita_screenshot" denseCaptions
```
<br />

When correctly set up, you should get a response similar this:
```json
{"modelVersion":"2023-10-01","denseCaptionsResult":{"values":[{"text":"a screenshot of a computer","confidence":0.8589847683906555,"boundingBox":{"x":0,"y":0,"w":2286,"h":1155}},{"text":"a blue dot on a black rectangle","confidence":0.6412748694419861,"boundingBox":{"x":771,"y":416,"w":862,"h":200}},{"text":"a screen shot of a whiteboard","confidence":0.7157526016235352,"boundingBox":{"x":666,"y":147,"w":1011,"h":587}},{"text":"a screenshot of a computer","confidence":0.8000496029853821,"boundingBox":{"x":197,"y":784,"w":1886,"h":348}},{"text":"a screenshot of a computer","confidence":0.843819260597229,"boundingBox":{"x":0,"y":15,"w":2258,"h":1089}},{"text":"a screenshot of a computer program","confidence":0.816832423210144,"boundingBox":{"x":1916,"y":814,"w":356,"h":293}},{"text":"a blue circle on a black stick","confidence":0.7007291316986084,"boundingBox":{"x":1186,"y":491,"w":79,"h":68}},{"text":"a screenshot of a computer","confidence":0.8584532141685486,"boundingBox":{"x":0,"y":109,"w":252,"h":721}}]},"metadata":{"width":2286,"height":1155}}
```
If you instead get an error like ``Failed to analyze image file '../barosa-screen-capture/Krita_screenshot'``, check your environment variables you set in the `.env` file for accuracy. This usually means you put your api key or endpoint incorrectly.
_It can also mean a variety of other things, but I recommend trying the basic troubleshoots first._

Once you're getting this json response, you're good to go on the Azure side.

<br /> 

### Webserver Environment Variables
This monolith uses a barebones `Express` NodeJS server for auth and for passing parameters to the front-end template for easy tweaking and debugging. _Yes, templating engines in 2024!_
I chose to use `ejs` for the front-end. _Personally, I couldn't find any reason to harbor a Next.JS React server-side rendering super ultra framework, just for rending a single html file with a few dynamic variables._

Open the `barosa-server` folder

And create another `.env` file, filling in the values below:
```env
BAROSA_BEARER_AUTH=make_up_a_password_to_here_to_replace_me
BAROSA_BEARER_SECRET=make_up_an_even_stronger_password_here
BAROSA_CLIENT_PORT=3000
```
`BAROSA_BEARER_AUTH` and `BAROSA_BEARER_SECRET` should be a long, unguessable phrase or password that is preferably more than 60 characters long. You can just smash your keyboard if you'd like

The reason we have these auth/secret vars is to protect you in case you decide to host this server publicly. The flow of how this works is:
- `BAROSA_BEARER_AUTH` is encrypted via JWT using `BAROSA_BEARER_SECRET` as the encryption key
- The encrypted auth is passed to the front-end template
- The front-end gives the encrypted auth to our GO backend, who decrypts it using the same `BAROSA_BEARER_SECRET` variable.

To test if you've set up the environment variables correctly for the GO server, do ``go run main.go`` inside of the `barosa-server` folder. You should see something like this:
```sh
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> barosa.fun/azure-ai-stream-backend/server.RequestPing (5 handlers)
[GIN-debug] GET    /image-features           --> barosa.fun/azure-ai-stream-backend/server.RequestImageFeatures (5 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```
If you see this, you're good to go!

**NOTE**: **Do not run the GO server like this!** this is just for testing to ensure the environment variables are set up properly

## Build
If you've followed the setup above with great success, then you can go ahead and build the backend!

Navigate to the root `backend` folder and you'll see a script called `build-go-modules`

This script runs every module's unit tests and compiles all of them to the `barosa-server` folder. 

The output of this command should be something like this:
```sh
 $ ./build-go-modules

------------------------------
superbarosa azure superplex
------------------------------
PASS
ok  	barosa.fun/azure-comms	0.002s
PASS
ok  	barosa.fun/capture	0.623s
?   	barosa.fun/shuferbarosa-ai	[no test files]
PASS
ok  	barosa.fun/azure-ai-stream-backend	0.077s
------------------------------
env file gossiped with frontend
Built 4 Go modules
```
If you see one or more unit tests failing: 
- Ensure you have `maim` and `xdotools` installed and accessible from your `$PATH`
- Check for any missing files. Make sure you're up to date with the latest `main` branch

If you _still_ have failing tests and you're certain you've set up correctly, [open an issue](https://github.com/shufsp/azure-ai-stream/issues/new)! I'm open to improvements within the scope of this project :)

## Run
_In progress_ of making a good way to run the entire project with one command. Right now as is, you must start the front-end and back-end separately.

### To start the front-end
Go into the `front-end` folder and run 
```sh
npm i && npm run build && npm run start
```
This installs the dependencies from `package.json`, builds to the `dist/` folder, and runs `dist/index.js`
You should now be able to visit the front end by going to ``http://localhost:3000`` in your browser.

### To start the back-end
Go into the `back-end` folder

Ensure everything is built by running:
```sh
./build-go-modules 
```
Ensure there are no failed unit tests or build errors, then you can start the server:
```sh
cd barosa-server && GIN_MODE=release ./barosa-server
```
This should output
```sh
2024/07/23 13:15:36 Loading env vars
2024/07/23 13:15:36 Starting barosa backend
```
Which means the server has started. You can now navigate to your front-end to test it out

## Help 
If you notice something wrong within the project, feel free to [open an issue](https://github.com/shufsp/azure-ai-stream/issues/new)
Even though I'm open to feedback and recommendations, I'd like to keep this project as simple as possible! Please keep any issues to bug fixes or platform support instead of giant feature requests


