pipeline {
    agent { 
        node {label 'slave-2'} 
    }

    stages {
        stage('Prepare job definition') {
            steps {
                echo 'Pulling...' + env.BRANCH_NAME
            }
            post {
                always { 
                    cleanWs()
                }
            }
        }
    }
}


