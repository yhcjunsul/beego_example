var loginApp = new Vue({
    el: '#login',
    data: {
        id: '',
        password: '',
        idLabel: 'ID',
        passwordLabel: 'Password',
        loginButtonLabel: 'Log In'
    },
    methods: {
        loginPost: function() {
            let form = new FormData()
            form.append('id', this.id)
            form.append('password', this.password)

            axios.post('http://localhost:8090/login', form).then(function(res) {
                alert("Log in 성공")
            }, function() {
                alert("Log in 실패")
            })
        }
    },
})