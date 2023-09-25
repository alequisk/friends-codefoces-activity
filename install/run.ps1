
$projectName = "main"
$processName = "$projectName.exe"

# Get the current script's directory
$currentScriptDirectory = Split-Path -Parent $MyInvocation.MyCommand.Path

# Set the project directory to the parent folder of the script
$projectDirectory = Join-Path -Path $currentScriptDirectory -ChildPath "..\"

# Check if Go is installed
$goInstalled = Get-Command go -ErrorAction SilentlyContinue

Write-Host "> Configuring to run 'Friends Codeforces Activity' app."

try {
  if ($args.Count -eq 0) {
    throw "Error: You must specify the handles that will be monitored after execute the executable: .\run.ps1 handle1 handle2"
  }

  # Set the working directory to the project directory
  Set-Location $projectDirectory
  if (-Not(Test-Path -Path $processName -PathType Leaf)) {
    if ($goInstalled) {
      Write-Host "> GoLang is installed on your system."
      Write-Host "> Building project..."
    
      # Build your Go project
      $buildOutput = go build .\cmd\main.go
    
      if ($LASTEXITCODE -eq 0) {
        Write-Host "> App build succeeded."
      }
      else {
        Write-Host "> App build failed." -ForegroundColor DarkRed
        throw $buildOutput
      }
    }
    else {
      throw "Error: GoLang is not installed on your system. Please install GoLang before building your project."
    }
  }

  # TODO: stop any other process if already running
  # Write-Host "Stoping run app if is running"
  # TASKKILL /IM $processName /F -ErrorAction SilentlyContinue

  Write-Output "> Putting app to run in background..."
  Start-Process $processName -WindowStyle Hidden -ArgumentList $args

  Write-Host "> Done! You will be notified when your friends submit a problem." -ForegroundColor DarkGreen
  Read-Host "Press enter to continue"
}
catch {
  Write-Host $_.Exception.Message -ForegroundColor Red
} 
