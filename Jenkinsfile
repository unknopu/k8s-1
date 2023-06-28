pipeline {
    agent { 
        node {label 'slave-2'} 
    }

    environment {
        BRANCH = "${env.BRANCH_NAME}"
    }

    stages {
        stage('Prepare job definition') {
            steps {
                script {
                    echo 'Pulling...' + env.BRANCH_NAME
                    withDockerRegistry(credentialsId: 'd2330c30-d21b-4670-84ed-37dd988b506a') {
                        sh '''
                            docker build --no-cache -t unknopu/${BRANCH} .
                            docker push unknopu/${BRANCH}
                        '''
                    }
                }
            }
            post {
                always { 
                    cleanWs()
                }
            }
        }
    }
}


