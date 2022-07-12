transaction() {
  prepare(signer: AuthAccount) {
    log("Preparing Hello World...")
  }
  execute {
    log("Hello World!!!!!")
  }
}