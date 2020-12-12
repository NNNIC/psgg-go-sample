# psgg-go-sample

I'd like to introduce how to use StateGo with Go programming language.

## Environment

Windows10  
[StateGo](https://statego.programanic.com)

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





 
 
statego golang 

Now developing.

referece
https://www.spinute.org/go-by-example/random-numbers.html

## ebiten with StateGo

developing at ./ebiten

### How to compile

<pre>
cd ebiten/rotatex3
go run main.go 
</pre>  

To view StateGo file (*.psgg) download and install ./statego/StateGo_0.67.20953.exe 

Play by clicking  
<a href="https://github.com/NNNIC/psgg-go-sample/blob/main/wiki/rotatex3.gif"><img src="https://raw.githubusercontent.com/NNNIC/psgg-go-sample/main/wiki/rotatex3.png" /></a>
