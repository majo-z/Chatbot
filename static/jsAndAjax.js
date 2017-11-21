//Researched and Adapted from:
//https://data-representation.github.io/notes/ajax.html
//https://api.jquery.com/keypress/
//https://api.jquery.com/event.preventdefault/
//https://msdn.microsoft.com/en-us/library/mt260494.aspx

const list = $(".list-group");// access list by class or id
//console.log(list);
const form = $("#userInput");
//console.log(form);

form.keypress(function(event){

    if(event.keyCode != 13){//13 - enter key
        return;// don't do anything unless enter key is pressed
    }

    //default behaviour of form is to make a new request - refreshes the page and clears everything
    //stop this, because we'll make requests with JS not form.
    event.preventDefault();

    const text = form.val(); //val() - get...hold onto it before it gets cleared
    form.val(""); //val() - set...set to empty string, 
    //console.log(text); //original text

    //if text is invalid - typed nothing, return
    if(!text.trim()){
        return;
    }

    //ajax - append new list item
    list.append("<li class=\"list-group-item\">" + text + "</li>");

    //.get makes jQuery do a get request
    $.get("/ask", {input:text}) //sends these parameters
        .done(function(resp){ //when it completes, this function executes, "resp" is the response from server
            list.append("<li class=\"list-group-item text-right\">" + resp + "</li>");//ajax - append new list item
        }).fail(function(){ //fail runs if anything goes wrong
            list.append("<li id=\"error\">CONNECTION TO ELIZA LOST!!</li>");
        });
});