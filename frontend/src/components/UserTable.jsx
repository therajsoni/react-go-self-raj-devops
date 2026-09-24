function UserTable({ users, onEdit, onDelete }) {
  if (users.length === 0) {
    return <div className="empty">No users found</div>;
  }

  return (
    <div className="table-wrapper">
      <table>
        <thead>
          <tr>
            <th>Photo</th>
            <th>Name</th>
            <th>Username</th>
            <th>Email</th>
            <th>Phone</th>
            <th>Joining Date</th>
            <th>Office Time</th>
            <th>Actions</th>
          </tr>
        </thead>

        <tbody>
          {users.map((user) => (
            <tr key={user._id}>
              <td>
                {user.photoBase64 ? (
                  <img
                    className="user-photo"
                    src={user.photoBase64}
                    alt={user.name}
                  />
                ) : (
                  <div className="no-photo">N/A</div>
                )}
              </td>

              <td>{user.name}</td>

              <td>{user.username}</td>

              <td>{user.email}</td>

              <td>{user.phoneNo}</td>

              <td>{user.joiningDate}</td>

              <td>
                {user.officeTimeIn} - {user.officeTimeOut}
              </td>

              <td>
                <div className="table-actions">
                  <button className="edit-btn" onClick={() => onEdit(user)}>
                    Edit
                  </button>

                  <button
                    className="delete-btn"
                    onClick={() => onDelete(user._id)}
                  >
                    Delete
                  </button>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default UserTable;
