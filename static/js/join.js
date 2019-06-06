var joinApp = new Vue({
    el: '#join',
    data: {
        id: '',
        password: '',
        name: '',
        idLabel: 'ID',
        passwordLabel: 'Password',
        nameLabel: 'Name',
        joinButtonLabel: 'Join'
    },
    methods: {
        moveToJoinPage: function() {
            window.location.href = 'http://localhost:8090/join'
        },
        joinPost: function() {
            let form = new FormData()
            form.append('id', this.id)
            form.append('password', this.password)
            form.append('name', this.name)

            axios.post('http://localhost:8090/member', form).then(function(res) {
                alert("회원가입 성공")
                window.location.href = 'http://localhost:8090'
            }, function() {
                alert("회원가입 실패")
            })
        }
    },
})