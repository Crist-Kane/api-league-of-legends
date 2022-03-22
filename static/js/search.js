let currentFilter = ""
function search(){
    let input = document.getElementById("input-search").value.toLowerCase();
    let list = document.getElementsByClassName("character-profile");
    
    for (let item of list) {
        if (!item.innerHTML.toLowerCase().includes(input)) {
            item.style.display="none";
        } else {
            item.style.display="block";         
        }
    }
}

function filter(newFilter){
    let container = document.getElementById("test");
    
    if (currentFilter == newFilter){
        for (let i = 0; i < container.children.length; i++){
            container.children[i].classList.remove("none");
        }
        currentFilter = "";
    }else{
        for (let i = 0; i < container.children.length; i++){
            const currElement = container.children[i];
            if (!currElement.classList.contains(newFilter)){
                currElement.classList.add("none");
            }else{
                currElement.classList.remove("none");
            }
        }
        currentFilter = newFilter;
    }
}

function assassin() { filter("Assassin"); }
function fighter() { filter("Fighter"); }
function mage() { filter("Mage"); }
function marksman() { filter("Marksman"); }
function support() { filter("Support"); }
function tank() { filter("Tank"); }