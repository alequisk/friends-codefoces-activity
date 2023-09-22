
$projectName = "main"
$processName = "$projectName.exe"

# Get the current script's directory
$currentScriptDirectory = Split-Path -Parent $MyInvocation.MyCommand.Path

# Set the project directory to the parent folder of the script
$projectDirectory = Join-Path -Path $currentScriptDirectory -ChildPath "..\"

# Check if Go is installed
$goInstalled = Get-Command go -ErrorAction SilentlyContinue

Write-Host "## Configuring to run FCA app." -ForegroundColor DarkGreen

try {
  # Set the working directory to the project directory
  Set-Location $projectDirectory
  if (-Not(Test-Path -Path $processName -PathType Leaf)) {
    if ($goInstalled) {
      Write-Host "- GOLang is installed on your system." -ForegroundColor DarkGreen
      Write-Host "- Building project..."
    
      # Build your Go project
      $buildOutput = go build .\cmd\main.go
    
      if ($LASTEXITCODE -eq 0) {
        Write-Host " - Go app build succeeded." -ForegroundColor DarkGreen
      }
      else {
        Write-Host " - Go app build failed." -ForegroundColor DarkRed
        throw $buildOutput
      }
    }
    else {
      throw "Go is not installed on your system. Please install Go before building your project."
    }
  }

  # TODO: stop any other process if already running
  # Write-Host "Stoping run app if is running"
  # TASKKILL /IM $processName /F -ErrorAction SilentlyContinue

  Write-Output "- Put app to run in background..."
  Start-Process $processName -WindowStyle Hidden

  Write-Host "- Done! You will be notified when your friends submit a problem." -ForegroundColor Green
  
}
catch {
  Write-Host $_.Exception.Message -ForegroundColor Red
} 

Read-Host "Press any key to continue"
