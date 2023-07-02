def buildApp() {
    echo "Build version: ${params.VERSION}"
}

def testApp() {
    echo 'testing the application ...'
}

def deployApp() {
    echo 'deploying the application ...'
}

return this