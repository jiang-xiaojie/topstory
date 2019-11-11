const current = 'dev'
const envs = {
  'dev': {
    'online': false,
    'domain': 'http://127.0.0.1:8765'
  },
  'test': {
    'online': false,
    'domain': 'http://127.0.0.1:8080'
  },
  'prod': {
    'online': true,
    'domain': 'https://api.xxx.com'
  }
}

const env = envs[current]
const online = env['online']
const domain = env['domain']

module.exports = {
  online,
  domain
}