async function main() {
    const { newEnforcer } = require('casbin');

    const enforcer = await newEnforcer('model.conf', 'policy.csv');

    console.log(await enforcer.enforce("user:default/guest", "orchestrator.workflows.foo", "read"));
    console.log(await enforcer.enforce("user:default/guest", "orchestrator.foo", "read"));
}

main();
