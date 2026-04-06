while(1){
    console.log(Date.now())
    await new Promise((res) => {
        setTimeout(res, 1000)
    })
}

export {}