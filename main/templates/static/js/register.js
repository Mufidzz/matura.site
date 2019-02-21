function submitForm(){
    var formData = {
            "ID": $("#id").val(),
            "PwordHash": $("#pw").val(),

            "Nama": $("#nama").val(),
            "POB": $("#tempat_lahir").val(),
            "DOB": $("#tanggal_lahir").val(),
            "TahunLulus": $("#tahun_lulus").val(),

            "AlamatA": $("#alamat_asal").val() + 
                    "," + $("#ac-kec-a").val() + 
                    "," + $("#ac-kab-a").val() + 
                    "," + $("#ac-sel-a").val(),
        
            "AlamatB": $("#alamat_tinggal").val() + 
                    "," + $("#ac-kec-b").val() + 
                    "," + $("#ac-kab-b").val() + 
                    "," + $("#ac-sel-b").val(),
        
            "Status": $("#status").val(),
            "Instansi": $("#instansi-sel").val() + 
                        "#" + $("#fakultas").val() + 
                        "#" + $("#prodi").val() + 
                        "#" + $("#instansi-txt").val(),
        
            "wa": $("#wa").val()
        };

    $.ajax({
      type: "POST",
      url: "http://192.168.137.1:1174/users",
      data: JSON.stringify(formData),
      success: function(){(data => 
                            {
                            Console.log (data);
                            }
                          )},
      dataType: "json",
      contentType : "application/json"
    });

    //window.location.replace("http://localhost:1174");
    console.log("Response : 200");
}

function getProvince(n){
    $.ajax({
      type: "VIEW",
      url: "http://localhost:1174/register/addr/province",
      data: { get_param: 'value' },
      success: function(jsonData){
                 cb = '';
                 $.each(jsonData, function(i,data){
                     cb+='<option value="'+data.id+'">'+data.provinsi.toUpperCase()+'</option>';
                 });
                 $(n).append(cb);
            },
      dataType: "json",
      contentType : "application/json"
    });
    
}

function getKabupaten(x,y){
    $.ajax({
      type: "VIEW",
      url: "http://localhost:1174/register/addr/kabupaten/" + $(x).val(),
      data: { get_param: 'value' },
      success: function(jsonData){
                 cb = '';
                 $.each(jsonData, function(i,data){
                     cb+='<option value="'+data.id+'">'+data.kabupaten.toUpperCase()+'</option>';
                 });
                 $(y).append(cb);
            },
      dataType: "json",
      contentType : "application/json"
    });
}

function getKecamatan(x,y){
    $.ajax({
      type: "VIEW",
      url: "http://localhost:1174/register/addr/kecamatan/" + $(x).val(),
      data: { get_param: 'value' },
      success: function(jsonData){
                 cb = '';
                 $.each(jsonData, function(i,data){
                     cb+='<option value="'+data.id+'">'+data.kecamatan.toUpperCase()+'</option>';
                 });
                 $(y).append(cb);
            },
      dataType: "json",
      contentType : "application/json"
    });
}

function tabAd(){
    if ( document.getElementById("status").value == 3){
        document.getElementById("status-detail").attribute = "disabled";
        document.getElementById("instansi-sel").style = "display: none;";
    }
}

function getUniversity(x){
    
    var i;
    
    $(x).val("");
    
    for(i=1; i < $(".ins-option").length ; i++){
        $(".ins-option").remove();
    }
    
    if($("#status-detail").val() != 99){
    
        $.ajax({
          type: "GET",
          url: "http://localhost:1174/addr/univ/" + $("#status-detail").val(),
          data: { get_param: 'value' },
          success: function(jsonData){
                     cb = '';
                     $.each(jsonData, function(i,data){
                         cb+='<option class="ins-option" value="'+data.instansi+'">'+data.instansi.toUpperCase()+'</option>'; 
                     });
                     $(x).append(cb);

                    console.log(data);
                },
          dataType: "json",
          contentType : "application/json"
        });
    }
}

function toggleStatusDetail(){
    
    if($("#status").val() != 1){
        $("#status-detail").attr("disabled", "disabled");
    } else {
        $("#status-detail").removeAttr("disabled");
    }
    
    if( $("#status").val() == 99 || $("#status").val() == 2 || $("#status-detail").val() == 99 ){
       $("#instansi-sel").attr("disabled", "disabled");
       $("#fakultas").attr("disabled", "disabled");
       $("#prodi").attr("disabled", "disabled");
       $("#instansi-txt").removeAttr("disabled");
       $("#instansi-txt").addClass(" ckr");
    } else {
       $("#instansi-txt").attr("disabled", "disabled");
       $("#instansi-sel").removeAttr("disabled");
       $("#fakultas").removeAttr("disabled");
       $("#prodi").removeAttr("disabled");
    }
}


function showTab(n) {
    
  var x = document.getElementsByClassName("tab");
  x[n].style.display = "block";
    
  if (n == 0) {
    document.getElementById("prevBtn").style.display = "none";
    document.getElementById("noBtn").style.display = "inline";
    document.getElementById("noBtn").style.opacity = "0";
  } else {
    document.getElementById("prevBtn").style.display = "inline";
      document.getElementById("noBtn").style.display = "none";
  }
  if (n == (x.length - 1)) {
    document.getElementById("nextBtn").innerHTML = "Finish";
  } else {
    document.getElementById("nextBtn").innerHTML = "Next";
  }
  fixStepIndicator(n)
}

function nextPrev(n) {
  var x = document.getElementsByClassName("tab");
    
  if (n == 1 && !validateForm()) return false;
  
  x[currentTab].style.display = "none";
 
  currentTab += n;
 
  if (currentTab >= x.length) {
    //document.getElementById("#regForm").submit();
    submitForm();
    return false;
  }
  showTab(currentTab);
}

function validateForm() {
  var x, y, i, valid = true;
  x = document.getElementsByClassName("tab");
  y = x[currentTab].getElementsByTagName("input");
  
  for (i = 0; i < y.length; i++) {
    if (y[i].value == "") {
      y[i].className += " invalid";
      valid = false;
    }
  }
  
  if (valid) {
    document.getElementsByClassName("step")[currentTab].className += " finish";
  }
  return valid; 
}


function fixStepIndicator(n) {
  var i, x = document.getElementsByClassName("step");
  for (i = 0; i < x.length; i++) {
    x[i].className = x[i].className.replace(" active", "");
  }
    
  x[n].className += " active";
}

function pwVisible(id) {
    var x = document.getElementById(id);
    if (x.type == "password") {
        x.type = "text";
    } else {
        x.type = "password";
    }
} 