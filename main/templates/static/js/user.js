function scroll() {
  if (window.pageYOffset >= 10) {
    stick.classList.add("sticky");
  } else {
    stick.classList.remove("sticky");
  }
}

function do_resize(textbox) {

 var maxrows=99; 
  var txt=textbox.value;
  var cols=70;

 var arraytxt=txt.split('\n');
  var rows=arraytxt.length; 

for (i=0;i<arraytxt.length;i++) 
  rows+=parseInt(arraytxt[i].length/cols);

 if (rows>maxrows) textbox.rows=maxrows;
  else textbox.rows=rows;
    
 $("#feed-status").scrollTop(9999);
 }

function autoResizer(){
    $("#feed-status").autoResize();
}