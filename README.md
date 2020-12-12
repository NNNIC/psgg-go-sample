# psgg-go-sample

I'd like to introduce how to use StateGo with Go programming language.

## Environment

Windows10  
[StateGo](https://statego.programanic.com)  
[Golang environment](https://golang.org/doc/install)  
[ebiten](https://ebiten.org/) (if you use ebiten framework.)

## Folder Tree

<pre>
--/  
  |   
  +- ebiten  : StateGo samples for 2D Framework Ebiten   
  |   
  +- general : a StateGo sample for general purposes.   
  |  
  +- test    : "try and error" tests to research.  
  |  
  +- test-ebiten : "try and error" ebiten tests to research.  
  |  
  +- wiki    :  etc.  
 </pre>
 
## A sample for general purposes.
Folder ``/general``  
Open ``testControl.psgg`` using StateGo.    
Without or with any framework, you may use this state machine as the template.
### State Chart
<img src="https://raw.githubusercontent.com/NNNIC/psgg-go-sample/main/wiki/g1.png" width=400px/>

### Run  
Open cmd.exe then type ...  
``
go run testControl.go
``  
![](https://raw.githubusercontent.com/NNNIC/psgg-go-sample/main/wiki/g1run.png)

## ebiten samples
Folder ``/ebiten``  

[ebiten](https://ebiten.org/) is a 2d game library.  
I created a sample framework to copolate with ebiten library. 

Draw apis must be called at draw function in "ebiten framework". So, I prepared drawlists for pushing draw apis at update timing. Then call drawlist at "ebiten's draw timing".

### sample #1
Folder ``/ebiten/sample1``

``mainCtonrol`` has a menu and five features. 
Four features are implementing simple features and ebiten's example.  
One feature is for calling general StateGo Control introduced above this section.

#### State Chart

``mainControl``  
<img src="https://raw.githubusercontent.com/NNNIC/psgg-go-sample/main/wiki/es1main.png" width=400px/>

``testControl`` (as is genaral's one)  
<img src="https://raw.githubusercontent.com/NNNIC/psgg-go-sample/main/wiki/es1test.png" width=400px/>

#### Run

![](https://raw.githubusercontent.com/NNNIC/psgg-go-sample/main/wiki/es1.gif)

### Sample #2

This sample only differs testControl from ''sample #1''

#### State Chart

``testControl`` 
<img src="https://raw.githubusercontent.com/NNNIC/psgg-go-sample/main/wiki/es2test.png" width=400px/>

### Playground
Folder ``/ebiten/playground``

This sample does not have ``testControl``.   
You may create and put ``testControl`` into this sample.  

### Gophers love ebiten
Folder ``/ebiten/gopherslove``

This a sample game using ebiten and StateGo.  

![](https://raw.githubusercontent.com/NNNIC/psgg-go-sample/main/wiki/rotatex3-2.gif)

### Webasembly version

[You can play it.](https://nnnic.github.io/psgg-go-sampleweb/)




