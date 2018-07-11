module.exports = function (context, myQueueItem) {
    context.log('JavaScript queue trigger function processed work item', myQueueItem);
    context.bindings.myOutputQueueItem = myQueueItem.toUpperCase()
    context.done();
};