function search(){
    let input = document.getElementById("input-search").value.toLowerCase();
    let list = document.getElementsByClassName("info");
    
    for (let item of list) {
        if (!item.innerHTML.toLowerCase().includes(input)) {
            item.style.display="none";
        } else {
            item.style.display="list-item";         
        }
    }
}