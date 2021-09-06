let btnTambah = document.getElementById("tambah-data-keluarga")
let btnKurang = document.getElementById("kurang-data-keluarga")
let familyContainer = document.getElementById("family")
let scriptBaris = familyContainer != null ?  [familyContainer.innerHTML] : []
if(btnTambah != null && btnKurang != null) {
    btnTambah.addEventListener("click", e => {
        e.preventDefault()
        scriptBaris.push(scriptBaris[0])
        familyContainer.innerHTML = scriptBaris.join(" ")
    })
    btnKurang.addEventListener("click", e => {
        e.preventDefault()
        scriptBaris.pop()
        let newBaris = scriptBaris.join(" ")
        familyContainer.innerHTML = newBaris
    })
}

let changer = document.querySelectorAll("#view-changer")
let tabs = document.querySelectorAll("#tab-view")
let identifier = null

if(changer.length > 0 && tabs.length > 0) {
    changer.forEach(c => {
        c.addEventListener("click", e => {
            identifier = e.target.dataset.identifier
            c.classList.add("btn-primary")
            changer.forEach(c2 => {
                if(c2.dataset.identifier != null && c2.dataset.identifier !== identifier) {
                    c2.classList.remove("btn-primary")
                } 
            });
            tabs.forEach(e => {
                e.dataset.identifier == identifier ? e.style.display = "block" : e.style.display = "none"
            });
        })
    })
}