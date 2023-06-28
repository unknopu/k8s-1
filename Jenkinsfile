pipeline {
    agent { 
        node {label 'slave-2'} 
    }

    stages {
        
        stage('Prepare job definition') {
            steps {
                script {
                    echo 'Pulling...' + env.BRANCH_NAME
                    withDockerRegistry(credentialsId: 'd2330c30-d21b-4670-84ed-37dd988b506a') {
                        sh '''
                            export IMAGE_NAME=jk$env.BRANCH_NAME
                            docker build --no-cache -t unknopu/$IMAGE_NAME .
                            docker push unknopu/$IMAGE_NAME
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


