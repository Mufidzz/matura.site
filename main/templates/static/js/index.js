function slide(n){
    showSlide(slideIndex +=n);
}
        
function showSlide(n){
    var i;
    var slides = document.getElementsByClassName("bg");
    var slidesb = document.getElementsByClassName("mid-txt");
    
    if (n > slides.length)
    { slideIndex = 1; }
    
    if (n < 1 )
    { slideIndex = slides.length; }
    
    for (i = 0; i < slides.length; i++) 
    { 
      slides[i].style.display = "none"; 
      slidesb[i].style.display = "none";
    }
    
    slides[slideIndex-1].style.display = "block";  
    slidesb[slideIndex-1].style.display = "block";  
}

function navbarFunc(minus = window.innerHeight - 200) {
  if (window.pageYOffset >=  minus) {
    navbar.classList.add("sticky");
  } else {
    navbar.classList.remove("sticky");
  }
}

function pageOneSlideController() {
  var i;
  var slides = document.getElementsByClassName("pageOneSlider");
  for (i = 0; i < slides.length; i++) {
    slides[i].style.display = "none";
  }
  pageOneSliderIndex++;
  if (pageOneSliderIndex > slides.length) {pageOneSliderIndex = 1}
  slides[pageOneSliderIndex-1].style.display = "block";
  setTimeout(pageOneSlideController, 3000);
} 