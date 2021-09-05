let btnTambah = document.getElementById("tambah-data-keluarga")
let btnKurang = document.getElementById("kurang-data-keluarga")
let familyContainer = document.getElementById("family")
let scriptBaris = [familyContainer.innerHTML]
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