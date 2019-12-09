export default {
  count: {
    subscribe(parent, args, { pubsub }, info) {
      let count = 0
      console.log(pubsub);

      setInterval(() => {
        count++
        pubsub.publish('count', {
          count
        })
      }, 1000)

      return pubsub.asyncIterator('count')
    }
  }
};
