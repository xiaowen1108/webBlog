var value=null,Base=null;
function getId(id){
	return document.getElementById(id);
};
function getTagName(tagname,id){
	return id?id.getElementsByTagName(tagname):document.getElementsByTagName(tagname);
};
function attr(obj,getAttr,setAttrValue){
	return setAttrValue?obj.setAttribute(getAttr,setAttrValue):obj.getAttribute(getAttr);
};
function getStyle(obj,attr){
	return obj.currentStyle?obj.currentStyle[attr]:getComputedStyle(obj)[attr];
};
function getClass(classStyle){
	var elements=getTagName('*');
	var arrClass=[];
	var str1='';
	for(var i=0;i<elements.length;i++){
		var reg=new RegExp(classStyle);
		if(elements[i].className.match(reg)){
			arrClass.push(elements[i]);
		};
	};
	return arrClass;
};
function action(obj,ev,fn){
	return obj.attachEvent?obj.attachEvent("on"+ev,fn):obj.addEventListener(ev,fn,false);
};
function shake(obj,attr,hz,endfn){
		if(obj.Shaketimer){return;};
		obj=getId(obj);
		var arr=[],
	    pos=parseInt(getStyle(obj,attr)),
	    posL=parseInt(getStyle(obj,'left')),
	    posT=parseInt(getStyle(obj,'top')),
	    num=0;
	    for(var i=hz;i>0;i-=2){
	    	arr.push(i,-i);
	    };
	    arr.push(0);
	    if(hz>40){
	    	msg('您设置的值: '+hz,' 抖动频率过大');
	    	return;
	    	};
	    obj.Shaketimer=setInterval(function(){
		if(attr==='left+top'||attr==='top+left'){
			obj.style.left=posL+arr[num]+'px';
			obj.style.top=posT+arr[num]+'px';
		}
		else{
			obj.style[attr]=pos+arr[num]+'px';
		}
		num++;
		if(num===arr.length){
			clearInterval(obj.Shaketimer);
			obj.Shaketimer=null;
			if(endfn){
				endfn();
			};
		};
	},50);
};
function animate(obj,attr,num,itarget,callBack){
	clearInterval(obj.timer);
	obj=getId(obj);
	num=parseInt(getStyle(obj,attr))<itarget?num:-num;
	obj.timer=setInterval(function(){
		var speed=parseInt(getStyle(obj,attr))+(num);
		if(speed>itarget&&num>0 || speed<itarget&&num<0){
			speed=itarget;
		};
		obj.style[attr]=speed+'px';
		if(speed==itarget){
			clearInterval(obj.timer);
			if(callBack){
				callBack();
			};
		};
	},15);
};
function tab(obj,on,Style){
	obj=getId(obj);
	var arrHead=[],
	    arrContent=[];
	var elements=getTagName('*',obj);
	for(var i=0;i<elements.length;i++){
		if(attr(elements[i],'tab')==='head'){
			arrHead.push(elements[i]);
		};
		if(attr(elements[i],'tab')==='content'){
			arrContent.push(elements[i]);
		};
	};
	for(var i=0;i<arrHead.length;i++){
		arrHead[i].index=i;
		action(arrHead[i],on,function(){
			for(var i=0;i<arrHead.length;i++){
				arrHead[i].className='';
				arrContent[i].style['display']='none';
			};
			this.className=Style;
			arrContent[this.index].style['display']='block';
		});
	};	
};
function placeHolder(obj,txt,style1,style2){
	obj=getId(obj);
	obj.value=txt;
	obj.className=style1;
	obj.onfocus=function(){
		if(this.value===txt){
			this.value='';
		    this.className='';
		}else{
			this.className=style2;
		};
	obj.onblur=function(){
		if(this.value.length===0){
			this.value=txt;
			this.className=style1;
		}else{
			this.style.color=style2;
		};
	};
	};
};
function enderNum(obj,obj2,MaxNum,left,top){
	obj=getId(obj);
	obj2=getId(obj2);
	left=parseInt(left);
	top=parseInt(top);
	obj.style.position='relative';
	var newContent=document.createElement('div'),
	    Style={
	    	'width':'100px',
	    	'color':'#666',
	    	'font-family':'Microsoft YaHei',
	    	'font-size':'12px',
	    	'position':'absolute',
	    	'left':left+'px',
	    	'top':top+'px',
			'z-index':'9999',
			'height':'30px'
	    };
	for(var i in Style){
		newContent.style[i]=Style[i];
	};
	if(MaxNum>800){
		msg('温馨提示！','字数过多!\n请输入符合规范的限制字数！');
		return;
	};
	obj.appendChild(newContent);
	newContent.innerHTML='还可以输入'+MaxNum+'字';
	obj2.onkeydown=function(){
		var arrTxt=[];
		var obj2Len=obj2.value.length;
		if(obj2Len>MaxNum){
			for(var j=0;j<obj2.value.length;j++){
				arrTxt.push(obj2.value.charAt(j));
			};
			arrTxt.pop();
			obj2.value=arrTxt.join('');
		}else{
			newContent.innerHTML='还可以输入'+(MaxNum-obj2Len)+'字';
		};
	};   
};
function getCode(){
	if(arguments.length>1)return;
	var arr=[arguments[0]];
	arr=arr.join('')
	var arrCode=[];
	for(var i=0;i<arr.length;i++){
		arrCode.push(arr.charAt(i)+'的Unicode码值为 '+arr.charCodeAt(i));
	};
	alert(arrCode);
	arrCode=[];//防止数组叠加缓存;
};

function dragElement(obj){
	var obj=getId(obj);
	num=1;
	obj.style.position='absolute';
	obj.style.zIndex='1';
	obj.onmouseover=function(){this.style.cursor='pointer';};
	obj.onmousedown=function(event)
	{
		this.style.cursor='pointer';
		this.style.zIndex=num++;
		var oevent=event||window.event;
		var disX=oevent.clientX-obj.offsetLeft;
		var disY=oevent.clientY-obj.offsetTop;
		document.onmousemove=function(event)
		{
			var oevent=event||window.event;
			var left=oevent.clientX-disX;
			var top=oevent.clientY-disY;
			if(left<0)left=0;
			if(top<0)top=0;
			obj.style.left=left+'px';
			obj.style.top=top+'px';
		};
		document.onmouseup=function()
		{
		document.onmousemove=null;
		document.onmouseup=null;
		};
		return false;
	};
};
function stopDrag(obj,boolean){
	if(boolean){
		obj=getTagName(obj);
		for(var i=0;i<obj.length;i++){
			obj[i].style['resize']='none';
		};
	}else{
		obj=getId(obj);
		obj.style['resize']='none';
	};
};
function transColor(obj,r1,g1,b1,r2,g2,b2){
	var num=1,
	    num2=1,
	    num3=1;
	timer=setInterval(function(){
		r1+=num+=5;
		g1+=num2+=5;
		b1+=num3+=5;
		document.title=r1+' '+g1+' '+b1
		if(r1>=r2){
			r1=r2;
		};
		if(g1>=g2){
			g1=g2;
		};
		if(b1>=b2){
			b1=b2;
		};
		obj.style.background='rgb('+r1+','+g1+','+b1+')';
		if(r1===r2&&g1===g2&&b1===b2){
			clearInterval(timer);
		};
	},20);	
};
function msg(title,content){
	title=!title?title='提示信息!':title=title;
	content=!content?content='<b>出现这句话通常是您没有填写任何有效的文本参数！</b>':content=content;
	getTagName('html')[0].style.height='100%';
	getTagName('body')[0].style.height='100%';
	var No_1=getTagName('*',document.body);
	var w_Height=(document.body.offsetHeight)/2-120;
	if(w_Height<0)w_Height=50;
	var father=document.createElement('div'),
	    message=document.createElement('div'),
	    head=document.createElement('div'),
	    Close=document.createElement('div'),
	    contain=document.createElement('p');
		message.id='message';
	var f_Style={
		'width':'100%',
		'height':'100%',
		'background':'transparent',
		'position':'fixed',
		'z-index':'99999',
		'font-family':'Microsoft YaHei,黑体'
	};
	var w_Style={
		'width':'300px',
		'height':'auto',
		'min-height':'120px',
		'background':'#fff',
		'box-shadow':'0px 0px 40px #888',
		'border-radius':'3px',
		'margin':'auto',
		'margin-top':w_Height+'px',
		'position':'relative',
		'overflow':'hidden'
	};
	var h_Style={
		'width':'100%',
		'height':'30px',
		'border-bottom':'1px solid #999',
		'text-indent':'1em',
		'font-size':'14px',
		'line-height':'30px',
		'background':'#ccc',
	};
	var c_Style={
		'color':'#fff',
		'text-align':'left',
		'font-size':'20px',
		'background':'#e94724',
		'width':'38px',
		'float':'right',
		'height':'31px'
	};
	var con_Style={
		'text-align':'left',
		'font-size':'13px',
		'padding':'0px 10px',
		'text-indent':'2em',
		'line-height':'18px',
		'text-shadow':'0px 0px 8px #ccc',
	};
	for(var i in f_Style){
		father.style[i]=f_Style[i];
	};
	for(var i in w_Style){
		message.style[i]=w_Style[i];
	};
	for(var i in h_Style){
		head.style[i]=h_Style[i];
	};
	for(var i in c_Style){
		Close.style[i]=c_Style[i];
	};
	for(var i in con_Style){
		contain.style[i]=con_Style[i];
	};
	head.innerHTML=title;
	Close.innerHTML='X';
	contain.innerHTML=content;
	document.body.insertBefore(father,No_1[0]);
	father.appendChild(message);
	message.appendChild(head);
	head.appendChild(Close);
	message.appendChild(contain);
	Close.onmouseover=function(){
		this.style.cursor='pointer';
		this.style.color='#000';
		};
	Close.onmouseout=function(){
		this.style.color='#fff';
		};
	Close.onclick=function(){
		message.style['minHeight']=null;
		animate('message','height',10,0,function(){
			document.body.removeChild(father);
		});
	};
};
function preview(){
	getTagName('html')[0].style.height='100%';
	getTagName('body')[0].style.height='100%';
	var No_1=getTagName('*',document.body);
	var elements=getTagName('img',document.body);
	var screenWidth=document.body.clientWidth;
	var screenHeight=document.body.clientHeight;
	var arrImg=[];
	var iframesFather=document.createElement('div'),
	    iframesImg=document.createElement('img'),
		iframesClose=document.createElement('div');
	    iframesImg.id='iframesImgX';
		iframesClose.id='iframesCloseX';
	var iF_Style={
		'background':'rgba(0,0,0,0.6)',
		'position':'fixed',
		'width':'100%',
		'height':'100%'
	},
	iImg_Style={
		'position':'absolute',
		'margin':'auto',
		'left':'0',
		'top':'0',
		'right':'0',
		'bottom':'0',
		'border':'3px double #fff'
		},
	iClose_Style={
		'width':'40px',
		'height':'40px',
		'position':'absolute',
		'color':'#000',
		'font-size':'25px',
		'line-height':'40px',
		'top':'15px',
		'right':'15px',
		'background':'#fff',
		'text-align':'center',
		'font-family':'Microsoft YaHei 黑体',
		'font-weight':'bold',
		'border-radius':'50%'
		};
	for(var i in iF_Style){
		iframesFather.style[i]=iF_Style[i];
		};
	for(var i in iImg_Style){
		iframesImg.style[i]=iImg_Style[i];
	};
	for(var i=0;i<elements.length;i++){
		if(attr(elements[i],'pre')==='yes'){
			arrImg.push(elements[i]);
		};
	};
	iframesClose.innerHTML='X';
	iframesClose.onmouseover=function(){
		this.style.cursor='pointer';
		};
	iframesClose.onclick=function(){
		animate('iframesCloseX','width',10,0,function(){
			iframesFather.removeChild(iframesClose);
			animate('iframesImgX','width',40,0,function(){
			document.body.removeChild(iframesFather);
		});
		});
	};
	for(var i=0;i<arrImg.length;i++){
		arrImg[i].onclick=function(){
			document.body.insertBefore(iframesFather,No_1[0]);
			iframesFather.appendChild(iframesClose);
			iframesFather.appendChild(iframesImg);
			for(var i in iClose_Style){
				iframesClose.style[i]=iClose_Style[i];
			};
			var nImg=new Image();
			    nImg.src=attr(this,'src'),
			    nHeight=nImg.height,
			    nWidth=nImg.width;
			if(nWidth>screenWidth||nHeight>screenHeight){
				if(parseInt(getStyle(iframesImg,'height'))>=screenHeight){
					iframesImg.style.height=(screenHeight-100)+'px';
				}else{
					if(screenWidth<=1300){
				iframesImg.src=attr(this,'src');
				iframesImg.style.width=(screenWidth-350)+'px';
						}else{
				iframesImg.src=attr(this,'src');
				iframesImg.style.width=(screenWidth-450)+'px';
						};
				};  
			}else{
			iframesImg.style.width=nWidth+'px';
			iframesImg.src=attr(this,'src');	
			};
		};
	};
};
function jumping(obj,url,seconds){
	seconds=!seconds?seconds=10:seconds;
	var numTime=seconds;
	var reg=/http:\/\//i;
	url='http://www.'+url;
	var obj=getId(obj);
	var timer=setInterval(function(){
		if(numTime>0){
		numTime--;
		obj.innerHTML='还有'+numTime+'秒,'+'您的浏览器即将跳转到这个网络地址: '+'<b style="color:red;">'+url.replace(reg,'')+'</b>';
			}else{
		clearInterval(timer);
		window.location=url;   
			};
		},1000);
};
function veryif(el){
	var obj=new Object();
	var t=true,
	    f=false;
	obj.email=function(){
		var reg=/^\w+@[a-z0-9]+\.[a-z]{2,4}$/;
		return (reg.test(el.value)?true:false);
	};
	obj.tel=function(){
		var reg=/^[1][3578]\d{9}$/;
		return (reg.test(el.value)?true:false);
	};
	obj.ip=function(){
		//var reg=//;
	};
	obj.cn=function(){
		var reg=/[\u4e00-\u9fa5]/;
		return (reg.test(el.value)?true:false);
	};
	obj.qq=function(){
		//var reg=//;
		return (reg.test(el.value)?true:false);
	};
	return obj;
};