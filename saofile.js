const spawn = require('cross-spawn')

module.exports = {
  description: 'Scaffolding out a go lang rpc module.',
  prompts() {
    return [
      {
        name: 'name',
        message: 'What is the name of the new project?',
        default: this.outFolder
      },
      {
        name: 'description',
        message: 'How would you descripe the new project?',
        default: `my go grpc project`
      },
      {
        name: 'author',
        message: 'What is your name',
        default: this.gitUser.name,
        store: true,
        required: true
      },
      {
        name: 'username',
        message: 'What is your GitHub username',
        default: this.gitUser.name,
        store: true,
        required: true
      },
      {
        name: 'email',
        message: 'What is your GitHub email',
        default: this.gitUser.email,
        store: true,
        required: true,
        validate: v => /.+@.+/.test(v)
      },
      {
        name: 'projectPath',
        message: 'Set your project path: ',
        required: true,
        // required
        validate: v => v !== ''
      },
      {
        name: 'website',
        default({ username }) {
          return `https://github.com/${username}`
        },
        store: true
      },
      {
        name: 'test',
        message: 'Use a ci?',
        type: 'confirm',
        default: true
      },
      {
        name: 'gateway',
        message: 'With grpc gateway?',
        type: 'confirm',
        default: true
      }
    ]
  },
  actions() {
    return [
      {
        type: 'add',
        files: '**',
        filters: {
          'circle.yml': 'test',
          'pb/pb.pb.gw.go': 'gateway',
          'pb/pb.swagger.json': 'gateway'
        }
      },
      {
        type: 'move',
        patterns: {
          // We keep `.gitignore` as `gitignore` in the project
          // Because when it's published to npm
          // `.gitignore` file will be ignored!
          gitignore: '.gitignore',
        }
      }
    ]
  },
  async completed() {
    spawn.sync('chmod', ['+x', './build.sh'], {
      stdio: 'ignore',
      cwd: this.outDir
    })
    this.gitInit()
    this.showProjectTips()
  }
}
