<container-list>

  <!-- layout -->
  <h3>{ opts.title }</h3>
  <form>
    <div class="list-options">
      <input onclick={ hiddenSwitch } checked={ this.showHidden } type="checkbox" id="confirmField">
      <label class="label-inline" for="confirmField">Show Exited</label>
    </div>
  </form>
  <table>
  <thead>
    <tr>
      <th>ID</th>
      <th>Status</th>
      <th>Command</th>
      <th>Created</th>
      <th>Image</th>
      <th>Names</th>
      <th>Ports</th>
    </tr>
  </thead>
  <tbody>
    <tr each={ container, i in activeContainers }>
      <td>{ container.ID }</td>
      <td>{ container.Status }</td>
      <td>{ container.Command }</td>
      <td>{ container.Created }</td>
      <td>{ container.Image.substring(0, 12) }</td>
      <td>{ container.Names }</td>
      <td>
        <ul each={ port in container.Ports }>
          <li>{port.Type} => {port.PrivatePort}</li>
        </ul>
      </td>
    </tr>
  </tbody>
</table>

  <!-- style -->
  <style scoped>
    h3 {
      font-size: 14px;
    }

    .list-options {
        float:right;
    }
  </style>

  <script>
    this.showHidden = false;

    hiddenSwitch = function(e){
      this.showHidden = !this.showHidden;
      this.activeContainers = filterContainers(opts.containers, this.showHidden);
    };

    filterContainers = function(containers, showHidden) {
      if (showHidden) {
          return containers;
      } else {
          return containers.filter((c) => {
            return c.Status.indexOf("Exited") === -1
          });
      }
    }

    this.activeContainers = filterContainers(opts.containers);

  </script>

</container-list>
