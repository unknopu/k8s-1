def gv 
pipeline {
    agent { 
        node {label 'slave-2'} 
    }
    parameters {
        // string(name: 'VERSION', defaultValue: '', description: 'version to deploy')
        choice(name: 'VERSION', choices: ['1.1.0', '1.2.0', '1.3.0'], description: '')
        booleanParam(name: 'EXECUTE_TESTS', defaultValue: true, description: '')
    }

    environment {
        BRANCH = "${env.BRANCH_NAME}"
        IMAGE_NAME = "app-${BRANCH}"
    }

    stages {
        stage('Init') {
            steps {
                script {
                    gv = load "script.groovy"
                }
            }
        }
        stage('Test') {
            // it will be run test if it is a dev or master.
            when {
                expression {
                    BRANCH == 'dev' || BRANCH == 'main' || params.EXECUTE_TESTS
                }
            }
            steps {
                gv.testApp()
            }
        }
        stage('Prepare job definition') {
            steps {
                script {
                    echo "Pulling... ${env.BRANCH_NAME}"
                    gv.buildApp()

                    withDockerRegistry(credentialsId: 'd2330c30-d21b-4670-84ed-37dd988b506a') {
                        sh '''
                            docker build --no-cache -t unknopu/${IMAGE_NAME} .
                            docker push unknopu/${IMAGE_NAME}
                        '''
                    }
                }
            }
            post {
                always { 
                    cleanWs()
                }
                success {
                    echo 'SUCCESS ...!!!'
                }
                failure {
                    echo 'ITS FAILED, NOOB !!!'
                }
            }
        }
    }
}


