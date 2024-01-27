# Session Manager

## Vision
For every new app or project that I have ever worked on, one of the
most common items which had missing elements or is present in every
app (I kid you not), is User management.
The field of User management is vast and diverse.

Few of the stories for User management can be categorized as such:
- Authentication methods
    - Session
    - Token
    - Oauth
- RBAC Protocols
    - LDAP
    - RADIUS
    - SAML
- Authorization
- Multi factor authentication
- Biometrics
- Auditing
- Time based access

Each of the above items can be further expanded into multi feature stories.

### So where do we start from?
Since this is a project not born out of need but out of enthusiasm to stop
writing the same code repetitively without actually solving the entire problem
in a wholesome manner, I want to start by focussing on providing an easy authentication API
layer dealing with token based authentication. Once that's built in, I want to
spend considerable effort in designing the authorization system.

### Problems and learnings from Authorization
Most products deal with authorization as a problem specific to their company.
They feel their needs are unique and cannot be solved in a wholesome manner.
Or maybe they don't want to spend enough time to solve it in a wholesome manner
because once it's built specific to their needs, it most probably won't change
for a long time. However, when you do have to change the requirements, it becomes
an uphill task to actually change it.
Other authorization procotols like LDAP have the concept of groups where a member
is attached to a group and then based on the inclusion in the group, a user has
access to a resource.

### Dynamic Groups
Dynamic groups are a way of including people in groups without actually having the
need to define the groups. Dynamic groups matches people based on certain criteria
that they match. For example, if you are building an executive report for the executives
of the company, the moment a person is promoted as an executive, the person should
become eligible for all executive reports. This sounds like a simple functionality
but the real challenge is doing it at scale for millions of users with hundreds of
data points.
