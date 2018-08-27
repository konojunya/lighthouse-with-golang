const input = require('fs')
  .readFileSync('/dev/stdin')
  .toString()

const stats = JSON.parse(input)
const scoreMap = Object.entries(stats.audits).reduce((acc, [key, a]) => {
  if (typeof a.score === 'number') {
    return Object.assign({}, acc, {[key]: a.score })
  }
  return acc
}, {})
console.log(JSON.stringify(scoreMap))
