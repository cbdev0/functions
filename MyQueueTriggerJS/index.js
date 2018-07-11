module.exports = function (context, myQueueItem) {
    context.log('JavaScript queue trigger function processed work item', myQueueItem);
    context.bindings["my-output-queue"] = myQueueItem.toUpperCase()
    context.done();
};